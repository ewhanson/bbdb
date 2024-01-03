<?php

namespace App\Console\Commands;

use App\Models\User;
use App\User\UserRoleEnum;
use Illuminate\Console\Command;

class SetUserRole extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'app:set-user-role {user_id}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Make user with given ID an admin';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        try {
            $userId = $this->argument('user_id');
            $user = User::find($userId);
            $user->role = UserRoleEnum::ADMIN;
            $user->save();
            $this->info("Updated user (ID: $userId)");
        } catch (\Exception $exception) {
            $this->error($exception->getMessage());
        }
    }
}
