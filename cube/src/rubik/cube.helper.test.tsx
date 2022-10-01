import { Colour, CreateCube, Cube, Side } from "./cube";
import { assertSameFaceColours } from "./face.test";

export function assertSameFaces (expected: Cube, actual: Cube) {
  assertSameFaceColours(expected.front, actual.front);
  assertSameFaceColours(expected.back, actual.back);
  assertSameFaceColours(expected.left, actual.left);
  assertSameFaceColours(expected.right, actual.right);
  assertSameFaceColours(expected.top, actual.top);
  assertSameFaceColours(expected.bottom, actual.bottom);
}

describe("Rotations Front", () => {
  test("Cube rotation clockwise front", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.front, false);

    cubeExpected.bottom.top.left.colour = Colour.Blue;
    cubeExpected.bottom.top.centre.colour = Colour.Blue;
    cubeExpected.bottom.top.right.colour = Colour.Blue;

    cubeExpected.left.bottom.right.colour = Colour.White;
    cubeExpected.left.centre.right.colour = Colour.White;
    cubeExpected.left.top.right.colour = Colour.White;

    cubeExpected.top.bottom.left.colour = Colour.Green;
    cubeExpected.top.bottom.centre.colour = Colour.Green;
    cubeExpected.top.bottom.right.colour = Colour.Green;

    cubeExpected.right.bottom.left.colour = Colour.Yellow;
    cubeExpected.right.centre.left.colour = Colour.Yellow;
    cubeExpected.right.top.left.colour = Colour.Yellow;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Cube rotation front counter clockwise", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.front, true);

    cubeExpected.bottom.top.left.colour = Colour.Green;
    cubeExpected.bottom.top.centre.colour = Colour.Green;
    cubeExpected.bottom.top.right.colour = Colour.Green;

    cubeExpected.left.bottom.right.colour = Colour.Yellow;
    cubeExpected.left.centre.right.colour = Colour.Yellow;
    cubeExpected.left.top.right.colour = Colour.Yellow;

    cubeExpected.top.bottom.left.colour = Colour.Blue;
    cubeExpected.top.bottom.centre.colour = Colour.Blue;
    cubeExpected.top.bottom.right.colour = Colour.Blue;

    cubeExpected.right.bottom.left.colour = Colour.White;
    cubeExpected.right.centre.left.colour = Colour.White;
    cubeExpected.right.top.left.colour = Colour.White;

    assertSameFaces(cubeExpected, cubeActual);
  });
});

describe("Rotations Back", () => {
  test("Cube rotation clockwise back", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.back, false);

    cubeExpected.bottom.bottom.left.colour = Colour.Green;
    cubeExpected.bottom.bottom.centre.colour = Colour.Green;
    cubeExpected.bottom.bottom.right.colour = Colour.Green;

    cubeExpected.left.bottom.left.colour = Colour.Yellow;
    cubeExpected.left.centre.left.colour = Colour.Yellow;
    cubeExpected.left.top.left.colour = Colour.Yellow;

    cubeExpected.top.top.left.colour = Colour.Blue;
    cubeExpected.top.top.centre.colour = Colour.Blue;
    cubeExpected.top.top.right.colour = Colour.Blue;

    cubeExpected.right.bottom.right.colour = Colour.White;
    cubeExpected.right.centre.right.colour = Colour.White;
    cubeExpected.right.top.right.colour = Colour.White;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Cube rotation back counter clockwise", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.back, true);

    cubeExpected.bottom.bottom.left.colour = Colour.Blue;
    cubeExpected.bottom.bottom.centre.colour = Colour.Blue;
    cubeExpected.bottom.bottom.right.colour = Colour.Blue;

    cubeExpected.left.bottom.left.colour = Colour.White;
    cubeExpected.left.centre.left.colour = Colour.White;
    cubeExpected.left.top.left.colour = Colour.White;

    cubeExpected.top.top.left.colour = Colour.Green;
    cubeExpected.top.top.centre.colour = Colour.Green;
    cubeExpected.top.top.right.colour = Colour.Green;

    cubeExpected.right.bottom.right.colour = Colour.Yellow;
    cubeExpected.right.centre.right.colour = Colour.Yellow;
    cubeExpected.right.top.right.colour = Colour.Yellow;

    assertSameFaces(cubeExpected, cubeActual);
  });
});

describe("Rotations left", () => {
  test("Cube rotation clockwise left", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.left, false);

    cubeExpected.bottom.bottom.left.colour = Colour.Orange;
    cubeExpected.bottom.centre.left.colour = Colour.Orange;
    cubeExpected.bottom.top.left.colour = Colour.Orange;

    cubeExpected.front.bottom.left.colour = Colour.Yellow;
    cubeExpected.front.centre.left.colour = Colour.Yellow;
    cubeExpected.front.top.left.colour = Colour.Yellow;

    cubeExpected.top.top.left.colour = Colour.Red;
    cubeExpected.top.centre.left.colour = Colour.Red;
    cubeExpected.top.bottom.left.colour = Colour.Red;

    cubeExpected.back.bottom.right.colour = Colour.White;
    cubeExpected.back.centre.right.colour = Colour.White;
    cubeExpected.back.top.right.colour = Colour.White;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Cube rotation left counter clockwise", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.left, true);

    cubeExpected.bottom.bottom.left.colour = Colour.Red;
    cubeExpected.bottom.centre.left.colour = Colour.Red;
    cubeExpected.bottom.top.left.colour = Colour.Red;

    cubeExpected.front.bottom.left.colour = Colour.White;
    cubeExpected.front.centre.left.colour = Colour.White;
    cubeExpected.front.top.left.colour = Colour.White;

    cubeExpected.top.top.left.colour = Colour.Orange;
    cubeExpected.top.centre.left.colour = Colour.Orange;
    cubeExpected.top.bottom.left.colour = Colour.Orange;

    cubeExpected.back.bottom.right.colour = Colour.Yellow;
    cubeExpected.back.centre.right.colour = Colour.Yellow;
    cubeExpected.back.top.right.colour = Colour.Yellow;

    assertSameFaces(cubeExpected, cubeActual);
  });
});

describe("Rotations right", () => {
  test("Cube rotation clockwise right", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.right, false);

    cubeExpected.bottom.bottom.right.colour = Colour.Red;
    cubeExpected.bottom.centre.right.colour = Colour.Red;
    cubeExpected.bottom.top.right.colour = Colour.Red;

    cubeExpected.front.bottom.right.colour = Colour.White;
    cubeExpected.front.centre.right.colour = Colour.White;
    cubeExpected.front.top.right.colour = Colour.White;

    cubeExpected.top.top.right.colour = Colour.Orange;
    cubeExpected.top.centre.right.colour = Colour.Orange;
    cubeExpected.top.bottom.right.colour = Colour.Orange;

    cubeExpected.back.bottom.left.colour = Colour.Yellow;
    cubeExpected.back.centre.left.colour = Colour.Yellow;
    cubeExpected.back.top.left.colour = Colour.Yellow;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Cube rotation right counter clockwise", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.right, true);

    cubeExpected.bottom.bottom.right.colour = Colour.Orange;
    cubeExpected.bottom.centre.right.colour = Colour.Orange;
    cubeExpected.bottom.top.right.colour = Colour.Orange;

    cubeExpected.front.bottom.right.colour = Colour.Yellow;
    cubeExpected.front.centre.right.colour = Colour.Yellow;
    cubeExpected.front.top.right.colour = Colour.Yellow;

    cubeExpected.top.top.right.colour = Colour.Red;
    cubeExpected.top.centre.right.colour = Colour.Red;
    cubeExpected.top.bottom.right.colour = Colour.Red;

    cubeExpected.back.bottom.left.colour = Colour.White;
    cubeExpected.back.centre.left.colour = Colour.White;
    cubeExpected.back.top.left.colour = Colour.White;

    assertSameFaces(cubeExpected, cubeActual);
  });
});

describe("Rotations bottom", () => {
  test("Cube rotation clockwise bottom", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.bottom, false);

    cubeExpected.front.bottom.left.colour = Colour.Green;
    cubeExpected.front.bottom.centre.colour = Colour.Green;
    cubeExpected.front.bottom.right.colour = Colour.Green;

    cubeExpected.back.bottom.left.colour = Colour.Blue;
    cubeExpected.back.bottom.centre.colour = Colour.Blue;
    cubeExpected.back.bottom.right.colour = Colour.Blue;

    cubeExpected.right.bottom.left.colour = Colour.Orange;
    cubeExpected.right.bottom.centre.colour = Colour.Orange;
    cubeExpected.right.bottom.right.colour = Colour.Orange;

    cubeExpected.left.bottom.left.colour = Colour.Red;
    cubeExpected.left.bottom.centre.colour = Colour.Red;
    cubeExpected.left.bottom.right.colour = Colour.Red;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Cube rotation bottom counter clockwise", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.bottom, true);

    cubeExpected.front.bottom.left.colour = Colour.Blue;
    cubeExpected.front.bottom.centre.colour = Colour.Blue;
    cubeExpected.front.bottom.right.colour = Colour.Blue;

    cubeExpected.back.bottom.left.colour = Colour.Green;
    cubeExpected.back.bottom.centre.colour = Colour.Green;
    cubeExpected.back.bottom.right.colour = Colour.Green;

    cubeExpected.right.bottom.left.colour = Colour.Red;
    cubeExpected.right.bottom.centre.colour = Colour.Red;
    cubeExpected.right.bottom.right.colour = Colour.Red;

    cubeExpected.left.bottom.left.colour = Colour.Orange;
    cubeExpected.left.bottom.centre.colour = Colour.Orange;
    cubeExpected.left.bottom.right.colour = Colour.Orange;

    assertSameFaces(cubeExpected, cubeActual);
  });
});

describe("Rotations top", () => {
  test("Cube rotation clockwise top", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.top, false);

    cubeExpected.front.top.left.colour = Colour.Blue;
    cubeExpected.front.top.centre.colour = Colour.Blue;
    cubeExpected.front.top.right.colour = Colour.Blue;

    cubeExpected.back.top.left.colour = Colour.Green;
    cubeExpected.back.top.centre.colour = Colour.Green;
    cubeExpected.back.top.right.colour = Colour.Green;

    cubeExpected.right.top.left.colour = Colour.Red;
    cubeExpected.right.top.centre.colour = Colour.Red;
    cubeExpected.right.top.right.colour = Colour.Red;

    cubeExpected.left.top.left.colour = Colour.Orange;
    cubeExpected.left.top.centre.colour = Colour.Orange;
    cubeExpected.left.top.right.colour = Colour.Orange;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Cube rotation top counter clockwise", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();
    cubeActual.RotateFace(Side.top, true);

    cubeExpected.front.top.left.colour = Colour.Green;
    cubeExpected.front.top.centre.colour = Colour.Green;
    cubeExpected.front.top.right.colour = Colour.Green;

    cubeExpected.back.top.left.colour = Colour.Blue;
    cubeExpected.back.top.centre.colour = Colour.Blue;
    cubeExpected.back.top.right.colour = Colour.Blue;

    cubeExpected.right.top.left.colour = Colour.Orange;
    cubeExpected.right.top.centre.colour = Colour.Orange;
    cubeExpected.right.top.right.colour = Colour.Orange;

    cubeExpected.left.top.left.colour = Colour.Red;
    cubeExpected.left.top.centre.colour = Colour.Red;
    cubeExpected.left.top.right.colour = Colour.Red;

    assertSameFaces(cubeExpected, cubeActual);
  });
});
