<?php

namespace App\Livewire;

use Illuminate\Support\Carbon;
use Livewire\Component;

class Footer extends Component
{
    private string $year;

    private string $version;

    public function mount()
    {
        $this->year = Carbon::now()->format('Y');
        // TODO: Make dynamic
        $this->version = '0.13.0';
    }

    public function render()
    {
        return view('livewire.footer')
            ->with(['year' => $this->year, 'version' => $this->version]);
    }
}
