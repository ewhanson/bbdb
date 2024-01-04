<x-main-content-layout>
    <h1 class="text-2xl font-bold mb-4">#{{ $tagName }}</h1>
    @if(count($posts) === 0)
        🤷 Oops... No photos found
    @endif
    @foreach($posts as $post)
        <x-photo-card wire:key="{{ $post->id }}" :post="$post"/>
    @endforeach
    @if(!$isLastPage)
        <div x-intersect.threshold.80="$wire.loadMore()">Scroll to load more 👇</div>
        {{--        <button wire:click="loadMore" class="btn btn-outline btn-sm">--}}
        {{--            Load more photos--}}
        {{--        </button>--}}
    @endif
    <div wire:loading wire:target="loadMore">
        <span class="loading loading-dots loading-sm"></span>
    </div>
</x-main-content-layout>

