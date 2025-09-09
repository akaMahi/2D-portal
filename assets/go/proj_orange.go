components {
  id: "projection"
  component: "/assets/scripts/projection.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"proj_orange\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/props.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.2
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"proj\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"circle\"\n"
  "  }\n"
  "  data: 7.5\n"
  "}\n"
  ""
}
embedded_components {
  id: "fac"
  type: "factory"
  data: "prototype: \"/assets/go/orange_portal.go\"\n"
  ""
}
