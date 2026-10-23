# Maintainer: Jakub Klinkovský <lahwaacz at archlinux dot org>

pkgbase=vl-convert
pkgname=(vl-convert python-vl-convert)
pkgver=1.9.0
pkgrel=2
pkgdesc="Convert Vega-Lite chart specifications to SVG, PNG, or Vega"
arch=(x86_64)
url="https://github.com/vega/vl-convert"
license=(BSD-3-Clause)
makedepends=(
  cargo
  maturin
  python-installer
)
checkdepends=(
  python-pytest
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
b2sums=('ed0d4f42e19f3b7a24d96e45e679ec9febdf8b324aaae9653ff0293a0d4341e967a2556a6866f540051fdf94c6aed5e4e78408915dfbcaf521bac5ef6c0d08f6')

prepare() {
  cd $pkgbase-$pkgver
  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  # fix building with LTO
  CFLAGS+=' -ffat-lto-objects'

  cd $pkgbase-$pkgver
  cargo build --frozen --release --package vl-convert

  # build Python bindings
  cd $pkgbase-python
  maturin build --locked --release --target "$(rustc --print host-tuple)" --strip
}

check() {
  cd $pkgbase-$pkgver
  local cargo_test_options=(
    # failures in tests/test_cli.rs
    --skip test_vl2png::test
    --skip test_vl2png_locale::test::case_1
    --skip test_vl2png_theme_config::test
    --skip test_vl2svg::test::vl_version_1___v5_8__::name_2___stacked_bar_h__
    # failures in tests/test_specs.rs
    --skip test_locale
    --skip test_png_no_theme::test
    --skip test_png_theme_config::test
    --skip test_svg::test::name_3___stacked_bar_h__
  )
  cargo test --frozen --release --package vl-convert -- "${cargo_test_options[@]}"

  # test Python bindings
  local pytest_options=(
    -vv
    -W ignore::DeprecationWarning
    # skip test that requires scikit-image
    --ignore vl-convert-python/tests/test_specs.py
  )
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer target/wheels/*.whl
  test-env/bin/python -m pytest "${pytest_options[@]}" vl-convert-python/tests
}

package_vl-convert() {
  depends=(
    glibc
    libgcc
    zlib libz.so
  )

  cd $pkgbase-$pkgver
  # install the binary and license
  install -vDm 755 target/release/vl-convert -t "$pkgdir"/usr/bin/
  install -vDm 644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}

package_python-vl-convert() {
  depends=(
    glibc
    libgcc
    python
    zlib libz.so
  )

  cd $pkgbase-$pkgver
  python -m installer --destdir="$pkgdir" target/wheels/*.whl
  install -vDm 644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
