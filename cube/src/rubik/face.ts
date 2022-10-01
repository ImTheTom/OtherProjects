import { Colour } from "./cube";
import { Row, CreateRowForColour } from "./row";

export class Face {
  top: Row;
  centre: Row;
  bottom: Row;

  constructor (top: Row, centre: Row, bottom: Row) {
    this.top = top;
    this.centre = centre;
    this.bottom = bottom;
  }

  Rotate (counterClockwise: boolean) {
    if (counterClockwise) {
      this.rotateCounterClockwise();
    } else {
      this.rotateClockwise();
    }
  }

  rotateClockwise () {
    const bottomRightTmp = this.bottom.right;
    this.bottom.right = this.top.right;
    this.top.right = this.top.left;
    this.top.left = this.bottom.left;
    this.bottom.left = bottomRightTmp;

    const bottomCentreTmp = this.bottom.centre;
    this.bottom.centre = this.centre.right;
    this.centre.right = this.top.centre;
    this.top.centre = this.centre.left;
    this.centre.left = bottomCentreTmp;
  }

  rotateCounterClockwise () {
    const bottomRightTmp = this.bottom.right;
    this.bottom.right = this.bottom.left;
    this.bottom.left = this.top.left;
    this.top.left = this.top.right;
    this.top.right = bottomRightTmp;

    const bottomCentreTmp = this.bottom.centre;
    this.bottom.centre = this.centre.left;
    this.centre.left = this.top.centre;
    this.top.centre = this.centre.right;
    this.centre.right = bottomCentreTmp;
  }

  IsSolved (): boolean {
    const colour = this.centre.centre.colour;
    if (!this.top.isRowSameColour(colour)) {
      return false;
    }
    if (!this.centre.isRowSameColour(colour)) {
      return false;
    }
    if (!this.bottom.isRowSameColour(colour)) {
      return false;
    }
    return true;
  }
};

export const CreateFaceForColour = (colour: Colour): Face => {
  return new Face(
    CreateRowForColour(colour),
    CreateRowForColour(colour),
    CreateRowForColour(colour)
  );
};
