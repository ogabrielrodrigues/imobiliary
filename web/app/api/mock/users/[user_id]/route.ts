import { users } from "@/mock/users";
import { NextRequest } from "next/server";

type UserIDRequest = {
  params: {
    user_id: string;
  };
};

export async function GET(request: NextRequest, { params }: UserIDRequest) {
  return Response.json(users.find(user => user.id === params.user_id))
}