<?php

describe('unauthenticated smoke test', function () {
    it('should render unauthenticated pages', function (string $path) {
        $response = $this->get($path);
        $response->assertStatus(200);
    })->with(['/', '/about', '/whats-new', 'login']);

    it('should redirect to login for auth pages', function (string $path) {
        $response = $this->get($path);
        $response->assertStatus(302);
    })->with(['/feed', '/signup']);
});

describe('authenticated smoke test', function () {
    todo('Create auth tests with test enviornment')->with(['/feed', '/signup']);
});
