<?php

namespace App\Models;

use App\Events\SubscriberCreated;
use Illuminate\Database\Eloquent\Concerns\HasUlids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

/**
 * @property string $id
 * @property string $name
 * @property string $email
 */
class Subscriber extends Model
{
    use HasFactory, HasUlids;

    protected $fillable = [
        'name',
        'email',
    ];

    protected $dispatchesEvents = [
        'created' => SubscriberCreated::class,
    ];
}
