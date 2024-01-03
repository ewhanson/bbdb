<x-main-content-layout>
    <div class="card bg-base-100 shadow-xl w-auto">
        <form class="card-body" wire:submit="save">
            <h2 class="card-title">Notifications Signup</h2>
            <p>
                Sign up to receive an email when there are new photos. <br/>
                We'll email you at most once a day.
            </p>
            <div class="form-control w-full max-w-lg">
                <label class="label">
                    <span class="label-text">First Name</span>
                </label>
                <input
                    type="text"
                    wire:model="name"
                    placeholder="Enter your first name"
                    class="input input-bordered w-full"
                    required
                />
            </div>
            <div class="form-control w-full max-w-lg">
                <label class="label">
                    <span class="label-text">Email</span>
                </label>
                <input
                    type="email"
                    wire:model="email"
                    placeholder="Enter your email"
                    class="input input-bordered w-full"
                    required
                />
            </div>
            <div class="card-actions justify-end">
                <button type="submit" class="btn btn-primary">
                    Submit
                    <span wire:loading class="loading loading-spinner"></span>
                </button>
            </div>
        </form>
    </div>
</x-main-content-layout>
