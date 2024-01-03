<?php

namespace App\Livewire\Pages;

use App\Models\Subscriber;
use Livewire\Attributes\Validate;
use Livewire\Component;

class NotificationsSignup extends Component
{
    public string $errorMessage = '';

    public string $successMessage = '';

    #[Validate('required')]
    public string $name = '';

    #[Validate('required|unique:subscribers,email')]
    public string $email = '';

    public function save()
    {
        try {
            $this->validate();

            Subscriber::create(
                $this->only(['name', 'email'])
            );

            session()->flash('successMessage', 'Sign up successful! Check your email for details.');
            $this->reset();
        } catch (\Throwable $exception) {
            session()->flash('errorMessage', $exception->getMessage());
        }
    }

    public function render()
    {
        return view('livewire.pages.notifications-signup');
    }
}
