<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/

Route::get('/', \App\Livewire\Pages\Landing::class)->middleware('guest')->name('landing');
Route::get('/login', \App\Livewire\Pages\Login::class)->name('login');
Route::get('/about', \App\Livewire\Pages\About::class)->name('about');
Route::get('whats-new', \App\Livewire\Pages\WhatsNew::class)->name('whats-new');
Route::get('/feed', \App\Livewire\Pages\Feed::class)->name('feed')->middleware('auth:web');
Route::get('photos/{post}', \App\Livewire\Pages\SinglePhoto::class)->name('single-photo')->middleware('auth:web');
Route::get('tags/{tag}', \App\Livewire\Pages\TagFeed::class)->name('tag')->middleware('auth:web');
