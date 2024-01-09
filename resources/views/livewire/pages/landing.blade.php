<div>
    <main class="hero bg-base-200" style="min-height: 92.5vh">
        <div class="hero-content text-center">
            <div class="max-w-md">
                <h1 class="text-4xl font-bold">Welcome to Babygramz!</h1>
                <p class="py-6">
                    A personal project for sharing baby photos with friends and
                    family. Already have an access code? Login by clicking the button
                    below!
                </p>
                <a href="{{ route('login') }}" wire:navigate class="btn btn-primary">
                    Start viewing
                </a>
            </div>
        </div>
    </main>
</div>
