<?php

namespace App\Jobs;

use App\Mail\UpdatesAvailable;
use App\Models\Post;
use App\Models\PostStatus;
use App\Models\Subscriber;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;
use Illuminate\Support\Facades\Mail;

class SendUpdateNotifications implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    /**
     * Create a new job instance.
     */
    public function __construct()
    {
        //
    }

    /**
     * Execute the job.
     */
    public function handle(): void
    {
        $posts = Post::whereHas('postStatus', function (Builder $query) {
            $query->where('notification_sent', '=', 0);
        })->get();

        $newPostCount = $posts->count();

        if ($newPostCount > 0) {
            Subscriber::all()->each(function ($subscriber) use ($newPostCount) {
                Mail::to($subscriber)->send(new UpdatesAvailable($subscriber, $newPostCount));
            });
        }

        PostStatus::where('notification_sent', '=', 0)
            ->update(['notification_sent' => 1]);
    }
}
