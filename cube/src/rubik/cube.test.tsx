import { Colour, CreateCube, Side } from "./cube";
import { assertSameFaces } from "./cube.helper.test";

test("Cube creation", () => {
  const cube = CreateCube();
  expect(cube.IsSolved()).toBeTruthy();
});

describe("Algorithm rotations", () => {
  // Used to solve first 2 layers
  test("Can do right trigger R U R'", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);

    cubeExpected.bottom.top.right.colour = Colour.Blue;

    cubeExpected.front.centre.right.colour = Colour.Yellow;
    cubeExpected.front.top.right.colour = Colour.Yellow;
    cubeExpected.front.top.centre.colour = Colour.Blue;
    cubeExpected.front.top.left.colour = Colour.Blue;

    cubeExpected.top.bottom.left.colour = Colour.Orange;
    cubeExpected.top.bottom.centre.colour = Colour.Orange;
    cubeExpected.top.bottom.right.colour = Colour.Green;

    cubeExpected.left.top.left.colour = Colour.Orange;
    cubeExpected.left.top.centre.colour = Colour.Orange;
    cubeExpected.left.top.right.colour = Colour.White;

    cubeExpected.right.top.left.colour = Colour.Red;
    cubeExpected.right.centre.left.colour = Colour.Red;
    cubeExpected.right.bottom.left.colour = Colour.Yellow;

    cubeExpected.back.top.centre.colour = Colour.Green;
    cubeExpected.back.top.right.colour = Colour.Green;

    assertSameFaces(cubeExpected, cubeActual);
  });

  // Used to solve first 2 layers
  test("Can do left trigger L' U' L", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.left, true);
    cubeActual.RotateFace(Side.top, true);
    cubeActual.RotateFace(Side.left, false);

    cubeExpected.bottom.top.left.colour = Colour.Green;

    cubeExpected.front.centre.left.colour = Colour.Yellow;
    cubeExpected.front.top.left.colour = Colour.Yellow;
    cubeExpected.front.top.centre.colour = Colour.Green;
    cubeExpected.front.top.right.colour = Colour.Green;

    cubeExpected.top.bottom.left.colour = Colour.Blue;
    cubeExpected.top.bottom.centre.colour = Colour.Orange;
    cubeExpected.top.bottom.right.colour = Colour.Orange;

    cubeExpected.left.top.right.colour = Colour.Red;
    cubeExpected.left.centre.right.colour = Colour.Red;
    cubeExpected.left.bottom.right.colour = Colour.Yellow;

    cubeExpected.right.top.left.colour = Colour.White;
    cubeExpected.right.top.centre.colour = Colour.Orange;
    cubeExpected.right.top.right.colour = Colour.Orange;

    cubeExpected.back.top.centre.colour = Colour.Blue;
    cubeExpected.back.top.left.colour = Colour.Blue;

    assertSameFaces(cubeExpected, cubeActual);
  });

  // Get yellow cross
  test("Can do F U R U' R' F'", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, true);
    cubeActual.RotateFace(Side.right, true);
    cubeActual.RotateFace(Side.front, true);

    cubeExpected.front.top.right.colour = Colour.Green;
    cubeExpected.front.top.centre.colour = Colour.Yellow;

    cubeExpected.right.top.left.colour = Colour.Orange;
    cubeExpected.right.top.centre.colour = Colour.Red;
    cubeExpected.right.top.right.colour = Colour.Red;

    cubeExpected.back.top.left.colour = Colour.Green;
    cubeExpected.back.top.centre.colour = Colour.Yellow;

    cubeExpected.left.top.left.colour = Colour.Yellow;
    cubeExpected.left.top.right.colour = Colour.Yellow;

    cubeExpected.top.top.centre.colour = Colour.Orange;
    cubeExpected.top.top.left.colour = Colour.Blue;
    cubeExpected.top.bottom.left.colour = Colour.Blue;
    cubeExpected.top.bottom.centre.colour = Colour.Blue;

    assertSameFaces(cubeExpected, cubeActual);
  });

  // Solve yellow face
  test("Can do R U R' U R U2 R'", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);

    cubeExpected.front.top.right.colour = Colour.Green;
    cubeExpected.front.top.left.colour = Colour.Yellow;

    cubeExpected.right.top.left.colour = Colour.Yellow;
    cubeExpected.right.top.centre.colour = Colour.Green;
    cubeExpected.right.top.right.colour = Colour.Green;

    cubeExpected.back.top.left.colour = Colour.Orange;
    cubeExpected.back.top.centre.colour = Colour.Blue;
    cubeExpected.back.top.right.colour = Colour.Blue;

    cubeExpected.left.top.left.colour = Colour.Yellow;
    cubeExpected.left.top.centre.colour = Colour.Red;
    cubeExpected.left.top.right.colour = Colour.Red;

    cubeExpected.top.bottom.left.colour = Colour.Blue;
    cubeExpected.top.bottom.right.colour = Colour.Red;
    cubeExpected.top.top.left.colour = Colour.Orange;

    assertSameFaces(cubeExpected, cubeActual);
  });

  // Solve the top corners
  test("Can do L' U R U' L U R' R U R' U R U2 R'", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.left, true);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, true);
    cubeActual.RotateFace(Side.left, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);

    cubeExpected.front.top.centre.colour = Colour.Blue;
    cubeExpected.front.top.right.colour = Colour.Blue;

    cubeExpected.right.top.left.colour = Colour.Red;
    cubeExpected.right.top.centre.colour = Colour.Orange;
    cubeExpected.right.top.right.colour = Colour.Orange;

    cubeExpected.back.top.left.colour = Colour.Blue;

    assertSameFaces(cubeExpected, cubeActual);
  });

  // Move edges clockwise
  test("Can do F2 U R' L F2 L' R U F2", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.right, true);
    cubeActual.RotateFace(Side.left, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.left, true);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.front, false);

    cubeExpected.front.top.centre.colour = Colour.Blue;
    cubeExpected.left.top.centre.colour = Colour.Orange;
    cubeExpected.right.top.centre.colour = Colour.Green;

    assertSameFaces(cubeExpected, cubeActual);
  });

  // Move edges counter clockwise
  test("Can do F2 U' R' L F2 L' R U' F2", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.top, true);
    cubeActual.RotateFace(Side.right, true);
    cubeActual.RotateFace(Side.left, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.left, true);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.top, true);
    cubeActual.RotateFace(Side.front, false);
    cubeActual.RotateFace(Side.front, false);

    cubeExpected.front.top.centre.colour = Colour.Green;
    cubeExpected.left.top.centre.colour = Colour.Blue;
    cubeExpected.right.top.centre.colour = Colour.Orange;

    assertSameFaces(cubeExpected, cubeActual);
  });
});

describe("Complex rotations", () => {
  test("Can use back rotations U B R B'", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.top, false);
    cubeActual.RotateFace(Side.back, false);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.back, true);

    cubeExpected.bottom.top.right.colour = Colour.Red;
    cubeExpected.bottom.centre.right.colour = Colour.Red;
    cubeExpected.bottom.bottom.right.colour = Colour.Red;
    cubeExpected.bottom.bottom.centre.colour = Colour.Red;

    cubeExpected.front.bottom.right.colour = Colour.Green;
    cubeExpected.front.centre.right.colour = Colour.White;
    cubeExpected.front.top.right.colour = Colour.White;
    cubeExpected.front.top.centre.colour = Colour.Blue;
    cubeExpected.front.top.left.colour = Colour.Blue;

    cubeExpected.left.top.left.colour = Colour.Orange;
    cubeExpected.left.top.centre.colour = Colour.Orange;
    cubeExpected.left.top.right.colour = Colour.Orange;
    cubeExpected.left.bottom.left.colour = Colour.Red;

    cubeExpected.right.bottom.left.colour = Colour.White;
    cubeExpected.right.bottom.centre.colour = Colour.White;
    cubeExpected.right.top.right.colour = Colour.Red;

    cubeExpected.back.top.left.colour = Colour.Green;
    cubeExpected.back.top.centre.colour = Colour.Green;
    cubeExpected.back.top.right.colour = Colour.Green;
    cubeExpected.back.bottom.left.colour = Colour.Yellow;
    cubeExpected.back.bottom.centre.colour = Colour.Yellow;
    cubeExpected.back.bottom.right.colour = Colour.Blue;

    cubeExpected.top.bottom.right.colour = Colour.Orange;
    cubeExpected.top.centre.right.colour = Colour.Orange;

    assertSameFaces(cubeExpected, cubeActual);
  });

  test("Can use down rotations L' B R B'", () => {
    const cubeExpected = CreateCube();
    const cubeActual = CreateCube();

    cubeActual.RotateFace(Side.left, true);
    cubeActual.RotateFace(Side.bottom, false);
    cubeActual.RotateFace(Side.right, false);
    cubeActual.RotateFace(Side.bottom, true);

    cubeExpected.bottom.top.right.colour = Colour.Red;
    cubeExpected.bottom.top.centre.colour = Colour.Red;
    cubeExpected.bottom.top.left.colour = Colour.Blue;
    cubeExpected.bottom.centre.left.colour = Colour.Red;
    cubeExpected.bottom.bottom.left.colour = Colour.Red;

    cubeExpected.front.bottom.centre.colour = Colour.Blue;
    cubeExpected.front.bottom.right.colour = Colour.Blue;
    cubeExpected.front.centre.left.colour = Colour.White;
    cubeExpected.front.centre.right.colour = Colour.White;
    cubeExpected.front.top.left.colour = Colour.White;
    cubeExpected.front.top.right.colour = Colour.Red;

    cubeExpected.top.bottom.left.colour = Colour.Orange;
    cubeExpected.top.bottom.right.colour = Colour.Green;
    cubeExpected.top.centre.left.colour = Colour.Orange;
    cubeExpected.top.centre.right.colour = Colour.Orange;
    cubeExpected.top.top.left.colour = Colour.Orange;
    cubeExpected.top.top.right.colour = Colour.Orange;

    cubeExpected.left.bottom.right.colour = Colour.White;

    cubeExpected.right.bottom.left.colour = Colour.Yellow;
    cubeExpected.right.centre.left.colour = Colour.Orange;
    cubeExpected.right.top.left.colour = Colour.White;

    cubeExpected.back.bottom.right.colour = Colour.Yellow;
    cubeExpected.back.centre.left.colour = Colour.Yellow;
    cubeExpected.back.centre.right.colour = Colour.Yellow;
    cubeExpected.back.top.left.colour = Colour.Yellow;
    cubeExpected.back.top.right.colour = Colour.Yellow;

    assertSameFaces(cubeExpected, cubeActual);
  });
});
