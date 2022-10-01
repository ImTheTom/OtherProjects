import React from "react";
import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders cube info", () => {
  render(<App />);
  const linkElement = screen.getByText(/Cube/i);
  expect(linkElement).toBeInTheDocument();
});
