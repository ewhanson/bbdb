<?php

namespace App\Livewire\Pages;

use App\Post\Post;
use Livewire\Component;

class Feed extends Component
{
    public array $posts;

    public int $page = 1;

    public bool $isLastPage = false;

    public function mount()
    {
        $results = Post::paginate(10, ['*'], 'page', $this->page);
        $this->isLastPage = $results->onLastPage();
        $this->posts = $results->items();
    }

    public function loadMore()
    {
        $this->page++;
        $results = Post::paginate(10, ['*'], 'page', $this->page);
        $this->isLastPage = $results->onLastPage();
        $this->posts = array_merge($this->posts, $results->items());
    }

    public function render()
    {
        return view('livewire.pages.feed');
    }
}
