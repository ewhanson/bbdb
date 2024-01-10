<?php

namespace App\Livewire;

use App\Livewire\Pages\Landing;
use App\Models\PostStatus;
use Carbon\Carbon;
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
                'hasRecentSiteUpdates' => $this->hasRecentSiteUpdates(),
            ]);
    }

    private function hasNewPhotos(): bool
    {
        return PostStatus::all()->count() > 0;
    }

    private function hasRecentSiteUpdates(): bool
    {
        return now()->subWeek()->lt(new Carbon(config('app.last_updated')));
    }
}
