<?php

namespace App\Post;

use Carbon\Carbon;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Spatie\Image\Manipulations;
use Spatie\MediaLibrary\HasMedia;
use Spatie\MediaLibrary\InteractsWithMedia;
use Spatie\MediaLibrary\MediaCollections\Models\Media;
use Spatie\Tags\HasTags;

/**
 * @property string $description
 * @property Carbon $date_taken
 * @property Carbon $created_at
 * @property Carbon $updated_at
 */
class Post extends Model implements HasMedia
{
    use HasFactory, HasTags, InteractsWithMedia;

    protected $fillable = [
        'description',
        'date_taken',
    ];

    protected $casts = [
        'date_taken' => 'date',
    ];

    // TODO: Consider removing 'preview' conversion and instead use reasonably sized responsive image instead
    public function registerMediaConversions(?Media $media = null): void
    {
        $this->addMediaConversion('preview')
            ->fit(Manipulations::FIT_CROP, 300, 300);
    }

    public function isNew(): bool
    {
        // TODO: Implement properly
        return true;
    }
}
