<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;
use App\Models\User;
use App\Models\Table;

class ReserveResource extends JsonResource
{
    public function toArray($request)
    {
        if ($request->get('user')) {
            $user = User::where('id', $request->get('user'))->firstOrFail();
        } else {
            $user = User::where('id', $this->user)->firstOrFail();
        }

        if ($request->get('table')) {
            $table = Table::where('id', $request->get('table'))->firstOrFail();
        } else {
            $table = Table::where('id', $this->table)->firstOrFail();
        }

        return [
            'id' => $this->id,
            'user' => $user,
            'table' => $table,
            'date' => $this->date,
            'turn' => $this->turn,
            'confirmed' => $this->confirmed,
        ];
    }
}
