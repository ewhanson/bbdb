<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="robots" content="noindex"/>
    <link rel="icon" type="image/svg+xml" href="/baby.svg"/>

    <title>{{ $title ?? 'Babygramz' }}</title>

    <!-- Styles -->
    @vite('resources/css/app.css')
</head>
<body>
<livewire:navbar/>
{{ $slot }}
<livewire:footer/>
</body>
</html>
