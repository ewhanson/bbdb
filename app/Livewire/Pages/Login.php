<?php

namespace App\Livewire\Pages;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Livewire\Attributes\Validate;
use Livewire\Component;

class Login extends Component
{
    #[Validate('required')]
    public string $password = '';

    public function login(Request $request): void
    {
        $this->validate();

        if (Auth::attempt([
            'email' => $this->password.'.user@babygramz.com',
            'password' => $this->password,
        ])) {
            $request->session()->regenerate();
            $this->redirect(Feed::class);
        } else {
            // TODO: Display error on what happened
        }
    }

    public function render()
    {
        return view('livewire.pages.login');
    }
}
