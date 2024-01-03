<x-mail::message>
    # Welcome, {{ $subscriber->name }}! 👋

    Thanks for signing up for Babygramz notifications.

    You will receive an email update whenever new photos are posted.

    <x-mail::footer>
        You are receiving this email because you subscribed to update notifications.

        [Unsubscribe.]({{ route('unsubscribe', ['id' => $subscriber->id]) }})
    </x-mail::footer>
</x-mail::message>
