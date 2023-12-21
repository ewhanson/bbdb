<?php

namespace App\Livewire;

use App\Livewire\Pages\Landing;
use Livewire\Component;

class Navbar extends Component
{
    public function logout(): void
    {
        auth()->logout();
        $this->redirect(Landing::class);
    }

    public function render()
    {
        return view('livewire.navbar')
            ->with([
                'hasNewPhotos' => $this->hasNewPhotos(),
                'isBuildOlderThanOneWeek' => $this->isBuildOlderThanOneWeek(),
            ]);
    }

    private function hasNewPhotos(): bool
    {
        // TODO: Implement
        return true;
    }

    private function isBuildOlderThanOneWeek(): bool
    {
        // TODO: Implement
        return false;
    }
}
