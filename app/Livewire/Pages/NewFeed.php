<?php

namespace App\Livewire\Pages;

use App\Models\Post;
use Livewire\Component;

class NewFeed extends Component
{
    /** @var Post[] */
    public array $posts;

    public int $page = 1;

    public bool $isLastPage = false;

    public function mount()
    {
        $results = Post::has('postStatus')->orderByDesc('date_taken')->paginate(10, ['*'], 'page', $this->page);
        $this->isLastPage = $results->onLastPage();
        $this->posts = $results->items();
    }

    public function loadMore()
    {
        $this->page++;
        $results = Post::has('postStatus')->orderByDesc('date_taken')->paginate(10, ['*'], 'page', $this->page);
        $this->isLastPage = $results->onLastPage();
        $this->posts = array_merge($this->posts, $results->items());
    }

    public function render()
    {
        return view('livewire.pages.new-feed');
    }
}
