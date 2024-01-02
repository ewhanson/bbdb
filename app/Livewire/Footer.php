<?php

namespace App\Livewire;

use Livewire\Component;

// TODO: Consider moving to plain Blade component
class Footer extends Component
{
    private string $version;

    public function mount()
    {
        // TODO: Make dynamic
        $this->version = '0.13.0';
    }

    public function render()
    {
        return view('livewire.footer')
            ->with(['version' => $this->version]);
    }
}
