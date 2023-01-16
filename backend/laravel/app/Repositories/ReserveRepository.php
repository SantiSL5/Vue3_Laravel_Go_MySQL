<?php

namespace App\Repositories;
use App\Models\Reserve;
use App\Interfaces\ReserveRepositoryInterface;
use App\Http\Requests\Reserve\CreateReserve;
use App\Http\Requests\Reserve\UpdateReserve;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class ReserveRepository implements ReserveRepositoryInterface
{

    public function getAllReservesRepo()
    {
        return Reserve::all();
    }

    public function getReserveByIdRepo($id) 
    {
        return Reserve::where('id', $id)->firstOrFail();
    }

    public function createReserveRepo(CreateReserve $request)
    {
        return Reserve::create($request->validated());
    }

    public function updateReserveRepo(UpdateReserve $request, $id)
    {
        return Reserve::where('id', $id)->update($request->validated());
    }

    public function deleteReserveRepo($id)
    {
        return Reserve::where('id', $id)->delete();
    }
}
