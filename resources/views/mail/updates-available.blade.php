<x-mail::message>
# Hey {{ $subscriber->name }},

Julian has {{ $newPostCount }} new {{ $photoNoun }} on Babygramz!

<x-mail::button :url="route('feed')">
View on Babygramz
</x-mail::button>

<x-mail::footer>
You are receiving this email because you subscribed to update notifications.

[Unsubscribe.]({{ $unsubscribeUrl }})
</x-mail::footer>
</x-mail::message>
