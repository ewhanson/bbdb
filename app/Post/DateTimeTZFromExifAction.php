<?php

namespace App\Post;

use Carbon\Carbon;
use Illuminate\Support\Facades\Log;

class DateTimeTZFromExifAction
{
    protected const DATE_TIME = 'DateTime';

    protected const DATE_TIME_ORIGINAL = 'DateTimeOriginal';

    protected const DATE_TIME_DIGITIZED = 'DateTimeDigitized';

    protected const OFFSET_TIME = 'UndefinedTag:0x9010';

    protected const OFFSET_TIME_ORIGINAL = 'UndefinedTag:0x9011';

    protected const OFFSET_TIME_DIGITIZED = 'UndefinedTag:0x9012';

    public function __construct(protected string $fullFilePath)
    {

    }

    public function execute(): ?Carbon
    {
        try {
            $exifResults = exif_read_data($this->fullFilePath, 'IFD0');

            if ($exifResults === false) {
                return null;
            }

            $dateTimeString = $this->getDateTime($exifResults);
            $timezoneOffsetString = $this->getTimezone($exifResults);

            if ($dateTimeString === null || $timezoneOffsetString === null) {
                return null;
            }

            $dateTime = Carbon::parseFromLocale($dateTimeString, null, $timezoneOffsetString);
            if ($dateTime !== null) {
                return $dateTime;
            }
        } catch (\Exception $exception) {
            Log::error("[Exif parsing error] : " . $exception->getMessage());
            return null;
        }

        return null;
    }

    private function getDateTime(array $exifResults): ?string
    {
        foreach ([self::DATE_TIME, self::DATE_TIME_ORIGINAL, self::DATE_TIME_DIGITIZED] as $key) {
            if (array_key_exists($key, $exifResults)) {
                return $exifResults[$key];
            }
        }

        return null;
    }

    private function getTimezone(array $exifResults): ?string
    {
        foreach ([self::OFFSET_TIME, self::OFFSET_TIME_ORIGINAL, self::OFFSET_TIME_DIGITIZED] as $key) {
            if (array_key_exists($key, $exifResults)) {
                return $exifResults[$key];
            }
        }

        return null;
    }
}
