<x-main-content-layout>
    <div class="card bg-base-100 shadow-xl w-auto sm:w-96">
        <form class="card-body" wire:submit="login">
            <h2 class="card-title">Babygramz Access</h2>
            <div class="form-control w-full max-w-md">
                <label class="label">
                    @error('password')
                    <span class="label-text text-error">
                        {{ $message }}
                    </span>
                    @else
                        <span class="label-text">
                        Please enter the password you were given for access.
                    </span>
                        @enderror
                </label>
                <input
                    wire:model="password"
                    type="password"
                    placeholder="Password"
                    class="input input-bordered w-full max-w-md @error('password') input-error @enderror"
                />
                <label class="label">
              <span class="label-text-alt">
                <a class="link" href="/admin">
                  Access uploader login
                </a>
              </span>
                </label>
            </div>
            <div class="card-actions justify-end">
                <button
                    type="submit"
                    class="btn btn-primary"
                >
                    Submit
                    <span wire:loading class="loading loading-spinner"></span>
                </button>
            </div>
        </form>
    </div>
</x-main-content-layout>
