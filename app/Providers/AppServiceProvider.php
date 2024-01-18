<?php

namespace App\Providers;

use Illuminate\Routing\UrlGenerator;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\ServiceProvider;
use Livewire\Livewire;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     */
    public function register(): void
    {
        //
    }

    /**
     * Bootstrap any application services.
     */
    public function boot(UrlGenerator $url): void
    {
        if (app()->environment('production')) {
            $url->forceScheme('https');
        }

        Livewire::setScriptRoute(function ($handle) {
            return Route::get('/vendor/livewire/livewire.js', $handle);
        });
    }
}
