# Fractage Documentation

## Endpoints

### Fractals

### Sierpinski Carpet

```powershell
curl localhost:6060/sierpinski-carpet
```

## Type Definitions

### Integer Type

**Format:** `(+|-)?[0-9]+`<br/>
**Definition:** A 64-bit signed integer.<br/>
**Alias:** `<int>`<br/>
**Example:** `768`

### Float Type

**Format:** `(+|-)?[0-9]+(.[0-9]+)?`<br/>
**Definition:** A 64-bit floating or fractional number.<br/>
**Alias:** `<float>`<br/>
**Example:** `7.68`

### Rectangle Type

#### Variant 1

**Format:** `<float>, <float>, <float>, <float>`<br/>
**Definition:** 4 comma-separated float values representing the _x_ positions, _y_ position, _width_ and _height_ of a rectangular area.<br/>
**Alias:** `<rect>`<br/>
**Example:** `1.13, 2, 9.8, 7`

#### Variant 2

**Format:** `<float>, <float>`<br/>
**Definition:** 2 comma-separated float values representing the _width_ and _height_ of a rectangular area. _x_ and _y_ values are 0.<br/>
**Alias:** `<rect>`<br/>
**Example:** `7.68, 7.86`
