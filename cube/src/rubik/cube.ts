import { rotateBack, rotateBottom, rotateFront, rotateLeft, rotateRight, rotateTop } from "./cube.helper";
import { CreateFaceForColour, Face } from "./face";

export const enum Colour {
  Unknown = 0,
  White,
  Green,
  Orange,
  Blue,
  Red,
  Yellow,
};

export const enum Side {
  Unknown = 0,
  back,
  left,
  right,
  front,
  bottom,
  top,
};

export class Cube {
  back: Face;
  left: Face;
  right: Face;
  front: Face;
  bottom: Face;
  top: Face;
  facing: Side;

  constructor (
    back: Face,
    left: Face,
    right: Face,
    front: Face,
    bottom: Face,
    top: Face
  ) {
    this.UpdateFaces(
      back,
      left,
      right,
      front,
      bottom,
      top
    );
    this.facing = Side.front;
  }

  UpdateFaces (
    back: Face,
    left: Face,
    right: Face,
    front: Face,
    bottom: Face,
    top: Face
  ) {
    this.back = back;
    this.left = left;
    this.right = right;
    this.front = front;
    this.bottom = bottom;
    this.top = top;
  }

  RotateFace (side: Side, counterClockwise: boolean = false) {
    switch (side) {
      case Side.left: {
        this.left.Rotate(counterClockwise);
        const newFaces = rotateLeft(this.top, this.bottom, this.front, this.back, counterClockwise);
        this.UpdateFaces(newFaces.back, this.left, this.right, newFaces.front, newFaces.bottom, newFaces.top);
        break;
      }
      case Side.right: {
        this.right.Rotate(counterClockwise);
        const newFaces = rotateRight(this.top, this.bottom, this.front, this.back, counterClockwise);
        this.UpdateFaces(newFaces.back, this.left, this.right, newFaces.front, newFaces.bottom, newFaces.top);
        break;
      }
      case Side.back: {
        this.back.Rotate(counterClockwise);
        const newFaces = rotateBack(this.top, this.bottom, this.left, this.right, counterClockwise);
        this.UpdateFaces(this.back, newFaces.left, newFaces.right, this.front, newFaces.bottom, newFaces.top);
        break;
      }
      case Side.front: {
        this.front.Rotate(counterClockwise);
        const newFaces = rotateFront(this.top, this.bottom, this.left, this.right, counterClockwise);
        this.UpdateFaces(this.back, newFaces.left, newFaces.right, this.front, newFaces.bottom, newFaces.top);
        break;
      }
      case Side.bottom: {
        this.bottom.Rotate(counterClockwise);
        const newFaces = rotateBottom(this.front, this.back, this.left, this.right, counterClockwise);
        this.UpdateFaces(newFaces.back, newFaces.left, newFaces.right, newFaces.front, this.bottom, this.top);
        break;
      }
      case Side.top: {
        this.top.Rotate(counterClockwise);
        const newFaces = rotateTop(this.front, this.back, this.left, this.right, counterClockwise);
        this.UpdateFaces(newFaces.back, newFaces.left, newFaces.right, newFaces.front, this.bottom, this.top);
        break;
      }
      default: {
        console.log("Unknown side");
      }
    }
  }

  IsSolved (): boolean {
    if (!this.back.IsSolved()) {
      return false;
    }
    if (!this.left.IsSolved()) {
      return false;
    }
    if (!this.right.IsSolved()) {
      return false;
    }
    if (!this.front.IsSolved()) {
      return false;
    }
    if (!this.bottom.IsSolved()) {
      return false;
    }
    if (!this.top.IsSolved()) {
      return false;
    }
    return true;
  }
};

export const CreateCube = (): Cube => {
  return new Cube(
    CreateFaceForColour(Colour.Red),
    CreateFaceForColour(Colour.Green),
    CreateFaceForColour(Colour.Blue),
    CreateFaceForColour(Colour.Orange),
    CreateFaceForColour(Colour.White),
    CreateFaceForColour(Colour.Yellow)
  );
};
