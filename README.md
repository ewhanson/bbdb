## Getting started

Babygramz is a digital scrapbook for sharing photos with friends and
family—a [home-cooked app](https://www.robinsloan.com/notes/home-cooked-app/). It allows you to create a
password-protected photo feed with the ethos of a physical scrapbook and feel of a modern social media app.

## Installation

> [!WARNING]
> Babygramz is mostly ready for others to test, but the use of the babygramz.com domain is still hardcoded in the
> application in a few places.

- Configure database settings in `.env`. I've tested it with `mariadb` and `sqlite`.
- Install PHP dependencies via `composer install`
- Install Node dependencies via `npm install && npm run build`
- Run migrations via `php artisan migrate`

## FAQ

<dl>
<dt>Why is the Github project named bbdb?</dt>
<dd>Because it's a DB for (photos of) BB 👶</dd>
</dl>

## TODOs

- The application was originally written using Pocketbase and Preact. Since I rewrote it using Laravel, I have not ported over the end-to-end tests nor fixed the Github action to properly run the PHP tests. 

