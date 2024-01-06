<div class="card card-compact bg-base-100 w-full sm:w-auto sm:max-w-md shadow-xl">
    {{-- TODO: See if custom class "min-w-28rem" is best approach --}}
    <figure class="sm:min-w-28rem">
        {{ $post->getFirstMedia()->img()->attributes(['id' => $uuid, 'class' => 'object-contain', 'alt' => "Shows $post->description"]) }}
    </figure>
    <div class="card-body">
        <div class="card-actions justify-end">
            <div title="{{ $altDate }}" class="badge badge-outline badge-sm">
                {{ $displayDate }}
            </div>
        </div>
        <h2 class="card-title">
            <a href="{{ route('single-photo', ['post' => $post->id]) }}" wire:navigate>{{ $post->description }}</a>
            @if ($post->isNew())
                <span class="badge badge-secondary">New</span>
            @endif
        </h2>
        <div class="card-actions">
            @foreach($post->tags as $tag)
                <a class="link link-hover" href="{{ route('tag', ['slug' => $tag->slug]) }}" wire:navigate>
                    #{{ $tag->name }}
                </a>
            @endforeach
        </div>
    </div>
</div>
</div>
