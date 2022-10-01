import { Colour } from "./cube";
import { Square } from "./square";

export class Row {
  left: Square;
  centre: Square;
  right: Square;

  constructor (left: Square, centre: Square, right: Square) {
    this.left = left;
    this.centre = centre;
    this.right = right;
  }

  isRowSameColour (colour: Colour): boolean {
    if (this.left.colour !== colour) {
      return false;
    }
    if (this.centre.colour !== colour) {
      return false;
    }
    if (this.right.colour !== colour) {
      return false;
    }
    return true;
  }
};

export const CreateRowForColour = (colour: Colour): Row => {
  return new Row(
    new Square(colour),
    new Square(colour),
    new Square(colour)
  );
};
