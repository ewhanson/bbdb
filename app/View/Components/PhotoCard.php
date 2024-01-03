<?php

namespace App\View\Components;

use App\Models\Post;
use Carbon\Carbon;
use Closure;
use Illuminate\Contracts\View\View;
use Illuminate\View\Component;
use Livewire\Attributes\Locked;

class PhotoCard extends Component
{
    #[Locked]
    public string $displayDate;

    #[Locked]
    public string $altDate;

    /**
     * Create a new component instance.
     */
    public function __construct(
        public Post $post
    ) {
        $this->altDate = $post->date_taken->format('Y-m-d, g:i a');
        if ($post->date_taken->lt(Carbon::now()->subWeek())) {
            $this->displayDate = $this->post->date_taken->format('M j, Y');
        } else {
            $this->displayDate = $post->date_taken->diffForHumans();
        }
    }

    /**
     * Get the view / contents that represent the component.
     */
    public function render(): View|Closure|string
    {
        return view('components.photo-card');
    }
}
