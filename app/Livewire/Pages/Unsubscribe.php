<?php

namespace App\Livewire\Pages;

use App\Models\Subscriber;
use Livewire\Component;

class Unsubscribe extends Component
{
    public function mount(string $id)
    {
        try {
            $subscriber = Subscriber::findOrFail($id);

            $subscriber->delete();
            session()->flash('successMessage', 'You have successfully unsubscribed from Babygramz notifications');
        } catch (\Throwable $exception) {
            session()->flash('errorMessage', 'This user has already unsubscribed or does not exist.');
        }
    }

    public function render()
    {
        return view('livewire.pages.unsubscribe');
    }
}
