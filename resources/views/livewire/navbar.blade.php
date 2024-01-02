<div class="navbar bg-base-100">
    <div class="flex-1">
        <a
            @auth
                href="/feed"
            @endauth
            @guest
                href="/"
            @endguest
            wire:navigate
            class="btn btn-ghost normal-case text-xl"
        >
            Babygramz👶🎆
        </a>
    </div>

    <div class="flex-none">
        <div class="dropdown dropdown-end">
            <div class="indicator">
                @if($hasNewPhotos)
                    <div class="badge badge-secondary badge-xs indicator-item mt-1 mr-1"></div>
                @endif
                <label tabIndex="0" class="btn btn-square btn-ghost">
                    <div>
                        <x-custom-icon type="horizontalDots"/>
                    </div>
                </label>
            </div>

            <ul
                tabIndex="0"
                class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52 z-50"
            >
                @auth
                    <li>
                        <a class="justify-between" href="/feed" wire:navigate>
                            Photo feed
                            @if($hasNewPhotos)
                                <span class="badge badge-secondary badge-sm">
                                     updated
                                </span>
                            @endif
                        </a>
                    </li>
                @endauth
                <li>
                    <a href="/about" wire:navigate>About</a>
                </li>
                <li>
                    <a href="/whats-new" wire:navigate class="justify-between">
                        What's new
                        @if($isBuildOlderThanOneWeek)
                            <span class="badge badge-secondary badge-sm">
                                updated
                              </span>
                        @endif
                    </a>
                </li>
                @auth
                    <li>
                        <a href="/signup" wire:navigate>
                            Notifications signup
                        </a>
                    </li>
                @endauth
                <li>
                    @auth
                        <button type="button" wire:click="logout">Logout</button>
                    @endauth
                    @guest
                        <a href="/login" wire:navigate>Login</a>
                    @endguest
                </li>
            </ul>
        </div>
    </div>
</div>
