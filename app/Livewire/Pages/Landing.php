<?php

namespace App\Livewire\Pages;

use Livewire\Component;

class Landing extends Component
{
    public function render()
    {
        return view('livewire.pages.landing')->with([
            'showFooter' => true,
        ]);
    }
}
