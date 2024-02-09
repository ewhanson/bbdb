<?php

namespace App\Listeners;

use App\Events\SubscriberCreated;
use App\Models\User;
use App\User\UserRoleEnum;
use Filament\Notifications\Notification;

class SendAdminNewSubscriberNotification
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
        $name = $event->subscriber->name;

        $recipients = User::where('role', '=', UserRoleEnum::ADMIN->value)->get();
        $recipients->each(function (User $recipient) use ($name) {
            Notification::make()
                ->title('New subscriber')
                ->info()
                ->icon('heroicon-o-user-plus')
                ->body("$name has signed up for email notifications.")
                ->sendToDatabase($recipient);
        });
    }
}
