<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class UserResource extends JsonResource
{
    public function toArray($request)
    {
        return [
            'id' => $this->id,
            'username' => $this->username,
            'email' => $this->email,
            'password' => $this->password,
            'type' => $this->type ? $this->type : "client",
            'image' => $this->image ? $this->image : "https://static.productionready.io/images/smiley-cyrus.jpg",
            'enabled' => $this->enabled ? $this->enabled : true,
        ];
    }
}
