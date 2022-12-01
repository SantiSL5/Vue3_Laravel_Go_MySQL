<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;
use App\Models\Category;

class TableResource extends JsonResource
{
    public function toArray($request)
    {
        if ($request->get('category')) {
            $category = Category::where('id', $request->get('category'))->firstOrFail();
        } else {
            $category = Category::where('id', $this->category)->firstOrFail();
        }

        return [
            'id' => $this->id,
            'code' => $this->code,
            'category' => $category,
            'capacity' => $this->capacity,
            'reserved' => $this->reserved,
        ];
    }
}
