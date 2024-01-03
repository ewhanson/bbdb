<?php

namespace App\Listeners;

use App\Events\SubscriberCreated;
use App\Mail\NotificationsSignup;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Support\Facades\Mail;

class SendNotificationsSignupEmail implements ShouldQueue
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
    public function handle(SubscriberCreated $event): void
    {
        $subscriber = $event->subscriber;

        Mail::to($subscriber)->send(new NotificationsSignup($subscriber));
    }
}
