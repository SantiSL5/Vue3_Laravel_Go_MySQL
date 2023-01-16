<?php

namespace App\Interfaces;
use App\Http\Requests\Reserve\CreateReserve;
use App\Http\Requests\Reserve\UpdateReserve;

interface ReserveRepositoryInterface
{
    function getAllReservesRepo();
    function getReserveByIdRepo($id);
    function createReserveRepo(CreateReserve $request);
    function updateReserveRepo(UpdateReserve $request,$id);
    function deleteReserveRepo($id);
}
