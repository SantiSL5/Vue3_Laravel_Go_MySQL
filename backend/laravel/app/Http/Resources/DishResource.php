<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;
use App\Models\DishType;

class DishResource extends JsonResource
{
    public function toArray($request)
    {
        if ($request->get('type')) {
            $type = DishType::where('id', $request->get('type'))->firstOrFail();
        } else {
            $type = DishType::where('id', $this->dishType)->firstOrFail();
        }

        return [
            'id' => $this->id,
            'dish' => $this->code,
            'type' => $type,
            'price' => $this->capacity,
            'desc' => $this->reserved,
            'photo' => $this->reserved,
        ];
    }
}
