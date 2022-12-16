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
            $type = DishType::where('id', $this->type)->firstOrFail();
        }

        return [
            'id' => $this->id,
            'dish' => $this->dish,
            'type' => $type,
            'price' => $this->price,
            'desc' => $this->desc,
            'photo' => $this->photo,
        ];
    }
}
