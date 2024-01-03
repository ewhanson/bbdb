<?php

namespace App\Listeners;

use App\Events\PostCreated;

class MakePostStatus
{
    /**
     * Create the event listener.
     */
    public function __construct()
    {
        //
    }

    /**
     * Handle the event.
     */
    public function handle(PostCreated $event): void
    {
        $event->post->postStatus()->create();
    }
}
