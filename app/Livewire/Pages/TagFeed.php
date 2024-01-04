<?php

namespace App\Livewire\Pages;

use App\Models\Post;
use Livewire\Component;
use Spatie\Tags\Tag;

class TagFeed extends Component
{
    /** @var Post[] */
    public array $posts;

    public string $tagName;

    public int $page = 1;

    public bool $isLastPage = false;

    public function mount(string $slug)
    {
        $tag = Tag::findFromString($slug);

        if ($tag === null) {
            $this->isLastPage = true;
            $this->posts = [];
            $this->tagName = '';

            return;
        }
        $results = Post::withAnyTags([$tag->name])->paginate(10, ['*'], 'page', $this->page);
        $this->isLastPage = $results->onLastPage();
        $this->posts = $results->items();
        $this->tagName = $tag->name;
    }

    public function loadMore()
    {
        $this->page++;
        $results = Post::withAnyTags([$this->tagName])->paginate(10, ['*'], 'page', $this->page);
        $this->isLastPage = $results->onLastPage();
        $this->posts = array_merge($this->posts, $results->items());
    }

    public function render()
    {
        return view('livewire.pages.tag-feed');
    }
}
