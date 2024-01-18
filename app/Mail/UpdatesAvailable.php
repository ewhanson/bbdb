<?php

namespace App\Mail;

use App\Models\Subscriber;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Mail\Mailable;
use Illuminate\Mail\Mailables\Content;
use Illuminate\Mail\Mailables\Envelope;
use Illuminate\Mail\Mailables\Headers;
use Illuminate\Queue\SerializesModels;

class UpdatesAvailable extends Mailable implements ShouldQueue
{
    use Queueable, SerializesModels;

    /**
     * Create a new message instance.
     */
    public function __construct(
        private Subscriber $subscriber,
        private int $newPostCount
    ) {
        //
    }

    /**
     * Get the message envelope.
     */
    public function envelope(): Envelope
    {
        return new Envelope(
            subject: 'Update: 📸 '.$this->newPostCount.' new '.$this->getPhotoNoun().' available',
        );
    }

    public function headers(): Headers
    {
        return new Headers(
            text: [
                'X-PM-Message-Stream' => 'outbound',
            ]
        );
    }

    /**
     * Get the message content definition.
     */
    public function content(): Content
    {
        return new Content(
            markdown: 'mail.updates-available',
            with: [
                'subscriber' => $this->subscriber,
                'newPostCount' => $this->newPostCount,
                'photoNoun' => $this->getPhotoNoun(),
                'unsubscribeUrl' => route('unsubscribe', ['id' => $this->subscriber->id]),
            ],
        );
    }

    /**
     * Get the attachments for the message.
     *
     * @return array<int, \Illuminate\Mail\Mailables\Attachment>
     */
    public function attachments(): array
    {
        return [];
    }

    private function getPhotoNoun(): string
    {
        if ($this->newPostCount > 1) {
            return 'photos';
        } else {
            return 'photo';
        }
    }
}
