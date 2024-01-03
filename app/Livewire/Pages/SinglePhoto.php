<?php

namespace App\Livewire\Pages;

use App\Models\Post;
use Livewire\Component;

class SinglePhoto extends Component
{
    public Post $post;

    public function render()
    {
        return view('livewire.pages.single-photo');
    }
}
