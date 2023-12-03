components {
  id: "camera"
  component: "/orthographic/camera.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "near_z"
    value: "-10.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "far_z"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "zoom"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "projection"
    value: "FIXED_ZOOM"
    type: PROPERTY_TYPE_HASH
  }
}
