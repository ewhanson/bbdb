<?php

namespace App\Console;

use App\Jobs\RemoveOutdatedPostStatuses;
use App\Jobs\SendUpdateNotifications;
use Illuminate\Console\Scheduling\Schedule;
use Illuminate\Foundation\Console\Kernel as ConsoleKernel;

class Kernel extends ConsoleKernel
{
    /**
     * Define the application's command schedule.
     */
    protected function schedule(Schedule $schedule): void
    {
        $schedule->job(new RemoveOutdatedPostStatuses())->dailyAt('07:00')->timezone('America/Vancouver');
        $schedule->job(new SendUpdateNotifications())->dailyAt('08:00')->timezone('America/Vancouver');
    }

    /**
     * Register the commands for the application.
     */
    protected function commands(): void
    {
        $this->load(__DIR__.'/Commands');

        require base_path('routes/console.php');
    }
}
