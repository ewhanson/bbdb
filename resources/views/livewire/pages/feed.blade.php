<x-main-content-layout>
    @if(count($posts) === 0)
        🤷 Oops... No photos found
    @endif
        @if($newPhotoCount > 0)
            <a href="{{ route('new-feed') }}" role="alert" class="alert w-full sm:max-w-md">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                     class="stroke-info shrink-0 w-6 h-6 hidden sm:block">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                <span>{{ $newPhotoCount }} new {{ $newPhotoCount === 1 ? 'photo' : 'photos' }}. Click to view.</span>
            </a>
        @endif
    @foreach($posts as $post)
            <x-photo-card wire:key="post-{{ $post->id }}" :post="$post"/>
    @endforeach
    @if(!$isLastPage)
            <div x-intersect.margin.500px="$wire.loadMore()">Scroll to load more 👇</div>
    @endif
    <div wire:loading wire:target="loadMore">
        <span class="loading loading-dots loading-sm"></span>
    </div>
</x-main-content-layout>

