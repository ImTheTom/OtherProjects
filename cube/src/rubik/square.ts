import { Colour } from "./cube";

export class Square {
  colour: Colour = Colour.White;

  constructor (colour: Colour) {
    this.colour = colour;
  };
};
