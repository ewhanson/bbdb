<x-main-content-layout>
    @if(count($posts) === 0)
        🤷 Oops... No photos found
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

