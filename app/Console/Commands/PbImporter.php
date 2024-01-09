<?php

namespace App\Console\Commands;

use App\Models\Post;
use App\Models\Subscriber;
use Carbon\Carbon;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\Http;

use function Laravel\Prompts\confirm;
use function Laravel\Prompts\info;
use function Laravel\Prompts\spin;

class PbImporter extends Command
{
    private const API_URL = 'https://babygramz.com/api';

    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'app:pb-importer {--count=300}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Imports all existing bbdb data (posts, tags, subscribers) from Pocketbase';

    private $apiToken = '';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        $confirmed = confirm('This should only be done once, bub. Are you sure you want to do this?');
        if (! $confirmed) {
            info('Good choice, boyo.');
        } else {
            info('Fine. Importing the stuff...');
            $this->apiToken = config('babygramz_token');

            $postCount = spin(
                fn () => $this->importPosts($this->option('count')),
                'Importing posts...'
            );
            $subCount = spin(
                fn () => $this->importSubscribers(),
                'Importing subscribers...'
            );

            info("Complete: $postCount post(s), $subCount subscriber(s) imported");

        }
    }

    private function importPosts(string $postCount = '300'): int
    {
        try {
            $response = Http::withToken($this->apiToken)
                ->get(self::API_URL.'/collections/photos/records?expand=tags&perPage='.$postCount.'&sort=dateTaken')
                ->json();
            $items = collect($response['items']);
            $results = $items->map(function ($item) {
                $dateTaken = (new Carbon($item['dateTaken']))->setTimezone('America/Vancouver');
                if (count($item['tags']) > 0) {
                    $tags = collect($item['expand']['tags'])->map(fn ($tag) => $tag['name'])->toarray();
                } else {
                    $tags = [];
                }
                $url = self::API_URL.'/files/'.$item['collectionId'].'/'.$item['id'].'/'.$item['file'];

                return [
                    'description' => $item['description'],
                    'date_taken' => $dateTaken,
                    'tags' => $tags,
                    'url' => $url,

                ];
            })
                ->each(function ($item) {
                    $post = new Post();

                    $post->date_taken = $item['date_taken'];
                    $post->description = $item['description'];
                    $post->saveQuietly();

                    if (! empty($item['tags'])) {
                        $post->attachTags($item['tags']);
                    }
                    $post->addMediaFromUrl($item['url'])
                        ->withResponsiveImages()
                        ->toMediaCollection();
                });

            return $results->count();
        } catch (\Exception $exception) {
            $this->error($exception->getMessage());
        }

        return 0;
    }

    private function importSubscribers(): int
    {
        try {
            $response = Http::withToken($this->apiToken)
                ->get(self::API_URL.'/collections/subscribers/records')
                ->json();
            $results = collect($response['items'])->map(function ($item) {
                return [
                    'name' => $item['name'],
                    'email' => $item['email'],
                ];
            })
                ->each(function ($item) {
                    try {
                        Subscriber::withoutEvents(fn () => Subscriber::create($item));
                    } catch (\Exception $exception) {
                        $this->error('Failed to import subscriber: '.$item['name'].'. Reason: '.$exception->getMessage());
                    }
                });

            return $results->count();
        } catch (\Exception $exception) {
            $this->error($exception->getMessage());
        }

        return 0;
    }
}
