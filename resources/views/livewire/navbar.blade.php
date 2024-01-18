<div class="navbar bg-base-100">
    <div class="flex-1">
        <a
            @auth
                href="{{ route('feed') }}"
            @endauth
            @guest
                href="{{ route('landing') }}"
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
                @auth()
                    @if(($hasNewPhotos || $hasRecentSiteUpdates))
                    <div class="badge badge-secondary badge-xs indicator-item mt-1 mr-1"></div>
                @endif
                @endauth
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
                        <a class="justify-between" href="{{ route('feed') }}" wire:navigate>
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
                    <a href="{{ route('about') }}" wire:navigate>About</a>
                </li>
                <li>
                    <a href="{{ route('whats-new') }}" wire:navigate class="justify-between">
                        What's new
                        @if($hasRecentSiteUpdates)
                            <span class="badge badge-secondary badge-sm">
                                updated
                              </span>
                        @endif
                    </a>
                </li>
                @auth
                    <li>
                        <a href="{{ route('signup') }}" wire:navigate>
                            Notifications signup
                        </a>
                    </li>
                        @if(auth()->user()->isPrivilegedUser())
                            <li>
                                <a href="/admin">Admin Panel</a>
                            </li>
                        @endif
                @endauth
                <li>
                    @auth
                        <button type="button" wire:click="logout">Logout</button>
                    @endauth
                    @guest
                            <a href="{{ route('login') }}" wire:navigate>Login</a>
                    @endguest
                </li>
            </ul>
        </div>
    </div>
</div>
