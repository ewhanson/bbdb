<x-main-content-layout>
    @foreach($posts as $post)
        <x-photo-card wire:key="{{ $post->id }}" :post="$post"/>
    @endforeach
    @if(!$isLastPage)
        <button wire:click="loadMore" class="btn btn-outline btn-sm">
            Load more photos
        </button>
    @endif
    <div wire:loading wire:target="loadMore">
        <span class="loading loading-dots loading-sm"></span>
    </div>
</x-main-content-layout>

