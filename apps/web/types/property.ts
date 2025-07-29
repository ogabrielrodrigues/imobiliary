import { Address } from "./address";

export type Property = {
  id: string
  water_id: string
  energy_id: string
  user_id: string
  owner_id: string
  kind: 'Residencial' | 'Comercial' | 'Industrial' | 'Terreno' | 'Rural',
  status: 'Disponível' | 'Ocupado' | 'Indisponível' | 'Reservado' | 'Reformando',
  address: Address,
};

